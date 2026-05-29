import { DailyRecord, PaymentStat, ApiData } from "./salesData";

interface AppStore {
    selectedDate: string;
    paymentStat: PaymentStat;
    data: DailyRecord;
    fetchData(api_url: string, date: string): Promise<ApiData | void>;
    saveRecord(dailyRecord: DailyRecord): Promise<ApiData | void>;
    init(): void;

}

export type { AppStore };