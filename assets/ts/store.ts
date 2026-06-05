import { raw, type Alpine } from 'alpinejs';
import type { PaymentStat, DailyRecord, ApiData } from './types/salesData';
import type { AppStore } from './types/stores';

declare global {
  interface Window {
    Alpine: Alpine;
  }
}

 const appStore: AppStore = {
    selectedDate: '',
    paymentStat: {} as PaymentStat,
    data: {} as DailyRecord,
    async fetchData(api_url: string, date: string): Promise<ApiData | void> {

        this.selectedDate = date;
        try {
            const response = await fetch(`${api_url}${date}`, { credentials: 'include' });
            if (!response.ok) throw new Error('ネットワーク応答が正常ではありませんでした');
            const rawData = await response.json() as ApiData;

            if(rawData) this.data = rawData.Record;
            // if (rawData.Found && rawData.Record.DateString !== date) {
            //     console.error('Fetched record date does not match requested date.');
            // }
            // if (!rawData.Found) rawData.Record.DateString = date;
            // this.data = rawData.Record;
            // if (!this.data.Found && this.data.Mode === '登録' && this.userId !== null) {
            //     this.data.Record.StaffCode = this.userId;
            // }
            console.log('called from form component', rawData);
            return rawData;
        } catch (e) {
            const err = e as Error;
            // this.error = `データの取得に失敗しました: ${err.message}`;
            console.error('Error fetching data:', e);
        } 
    },

    async saveRecord(dailyRecord: DailyRecord): Promise<ApiData> {
        dailyRecord.DateString = this.selectedDate;

        const recordToSend = JSON.parse(JSON.stringify(dailyRecord)) as DailyRecord;
        const response = await fetch('/api/sales-data', {
            method:      'POST',
            credentials: 'include',
            headers:     { 'Content-Type': 'application/json' },
            body:        JSON.stringify({ Record: recordToSend, Found: this.data!.Found }),
        });
        if (!response.ok) {
            throw new Error(`サーバーエラー: ${response.status} ${response.statusText}`);
        }
        const responseData = await response.json() as ApiData;
        this.data = responseData.Record;
        return responseData;
    },

    init() {
        // Initialize with default values or fetch initial data if needed
        console.log('App store initialized');
    }
};

document.addEventListener('alpine:init', () => {

    const Alpine = window.Alpine;
    Alpine.store('app', appStore);
 
});