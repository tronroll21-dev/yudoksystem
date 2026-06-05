import type { Alpine } from 'alpinejs';
import type { RearegiData, RearegiDetail } from './types/rearegiData';
import { AppStore } from './types/stores';

declare global {
  interface Window {
    Alpine: Alpine;
  }
}

interface Toast {
    show: boolean;
    message: string;
    type: 'success' | 'error';
}

enum RegiCsv {
  ShohinKubunCode = 1,
  ShohinKubunName = 2,
  Kingaku = 3,
  Suryo = 4,
  HeikinTanka = 5,
}


function parseField(items: string[], index: RegiCsv): string {
  return items[index].replace(/"/g, "").trim();
}

function parseLong(items: string[], index: RegiCsv): number {
  return parseInt(parseField(items, index), 10);
}

function computeRank(lRank: number): number {
  const tier = Math.trunc(lRank / 100);
  if (tier < 1 || tier > 5) return 100;
  return tier * 100;
}
    
async function submitRearegiDetail(hiduke: string, details: RearegiDetail[]): Promise<void> {
  // Placeholder for actual submission logic (e.g., API call)
  console.log("Submitting detail:", details);

 const api_url = '/api/rearegi-details';

    try {
        const response = await fetch(`${api_url}`, {
                method:      'POST',
                credentials: 'include',
                headers:     { 'Content-Type': 'application/json' },
                body:        JSON.stringify({ hiduke, records: details }),
            });
        if (!response.ok) throw new Error('ネットワーク応答が正常ではありませんでした');

        window.open(`/api/rearegi-report?start_date=${hiduke}&end_date=${hiduke}`, '_blank');

        return;
    } catch (e) {
        const err = e as Error;
        console.error('Error fetching data:', e);
    } 
}

const processFileOnFrontend = async (file: File, hiduke: Date): Promise<RearegiData> => {

    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = () => {
            const text = reader.result as string;
            const lines = text.split(/\r?\n/);

            // ── Accumulators ─────────────────────────────────────────────────────────
            let lngAmountKaisuuken = 0;
            let lngAmountMassage = 0;
            let lng6KaishukenSellCount = 0;
            let lngKaishukenSellCount = 0;

            const hidukeStr = hiduke.toISOString().slice(0, 10); // 'YYYY-MM-DD'
            const rearegiDetails: RearegiDetail[] = [];

            for (let lineNo = 0; lineNo < lines.length; lineNo++) {
                const strLine = lines[lineNo].trim();
                if (!strLine) continue;

                // Skip header row
                if (lineNo === 0) continue;

                const processRow = (): void => {
                    const items = strLine.split(",");

                    // Expect exactly 6 fields (indices 0–5)
                    if (items.length - 1 !== 5) return;

                    const lRank = parseLong(items, RegiCsv.ShohinKubunCode);
                    const lngKingaku = parseLong(items, RegiCsv.Kingaku);
                    const suryo = parseLong(items, RegiCsv.Suryo);

                    // Ticket count tracking
                    if (lRank === 165) lng6KaishukenSellCount = suryo;
                    if (lRank === 166) lngKaishukenSellCount = suryo;
                    if (lRank === 3) lng6KaishukenSellCount += suryo;
                    if (lRank === 4) lngKaishukenSellCount += suryo;

                    // Amount bucketing
                    if (lRank < 200 || lRank > 500) {
                    lngAmountKaisuuken += lngKingaku;
                    } else {
                    lngAmountMassage += lngKingaku;
                    }

                    // Queue the INSERT (executed below, inside the loop)
                    const rank = computeRank(lRank);
                    const shohinCode = parseLong(items, RegiCsv.ShohinKubunCode);
                    const shohinName = parseField(items, RegiCsv.ShohinKubunName);
                    const heikinTanka = parseLong(items, RegiCsv.HeikinTanka);

                    rearegiDetails.push({ hiduke: hidukeStr, rank, shohinCode, shohinName, kingaku: lngKingaku, suryo, heikinTanka });
                };
                processRow();
            }

            resolve({ 
        lngAmountKaisuuken,
        lngAmountMassage,
        lngKaishukenSellCount,
        lng6KaishukenSellCount,
         rearegiDetails });
        };
        reader.onerror = () => {
            reject(new Error('Error reading file'));
        };
        reader.readAsText(file, "Shift_JIS");
    });
};

document.addEventListener('alpine:init', () => {

    const Alpine = window.Alpine;
    Alpine.data('rearegiUpload', () => ({
        rearegiLocalData: {
            lngAmountKaisuuken: 0,
            lngAmountMassage: 0,
        } as RearegiData,
        formData: {
            files: null as FileList | null,
        },
        handleDropRearegi(event: DragEvent) {
            this.formData.files = event.dataTransfer!.files;
            this.processRearegiData();
        },
        async processRearegiData(): Promise<void> {
            
            const files = this.formData.files;
            if (!files || files.length === 0) {
                console.error('No file selected for processing.');
                this.$dispatch('toast', { show: true, message: 'ファイルが選択されていません', type: 'error' } as Toast);
                return;
            }


            const { lngAmountKaisuuken, lngAmountMassage, rearegiDetails } = await processFileOnFrontend(this.formData.files![0], new Date())  // ← passed in
            this.rearegiLocalData.lngAmountKaisuuken = lngAmountKaisuuken;
            this.rearegiLocalData.lngAmountMassage = lngAmountMassage;

            submitRearegiDetail((this.$store.app as AppStore).selectedDate, rearegiDetails);
        }
    }));

});