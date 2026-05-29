type RearegiDetail = {hiduke: string, rank: number, shohinCode: number, shohinName: string, kingaku: number, suryo: number, heikinTanka: number};

interface RearegiData {
    lngAmountKaisuuken: number;
    lngAmountMassage: number;
    lngKaishukenSellCount: number;
    lng6KaishukenSellCount: number;
    rearegiDetails: RearegiDetail[];
}

export type { RearegiData, RearegiDetail };