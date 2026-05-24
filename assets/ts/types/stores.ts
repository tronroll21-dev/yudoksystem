import { DailyRecord, PaymentStat, ApiData } from "./salesData";

interface AppStore {
    currentDate: string;
    paymentStat: PaymentStat;
    data: DailyRecord;
    fetchData(api_url: string, date: string): Promise<ApiData | void>;
    init(): void;
}

export type { AppStore };