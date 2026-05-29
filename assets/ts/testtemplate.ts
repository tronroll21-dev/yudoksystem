type Machine = "Machine1" | "Machine2" | "Machine3" | "Machine4" | "Machine5";
type PaymentMethod = "Cash" | "Denshi" | "Credit";
type Metric = "Count" | "Amount";

type MachinePaymentMetricKey = `${Machine}${PaymentMethod}${Metric}`;

type MachineReport = Partial<Record<MachinePaymentMetricKey, number>>;

const machineReport: MachineReport = {
    Machine1CashCount: 10,
    Machine1CashAmount: 1000,
    Machine1DenshiCount: 5,
    Machine1DenshiAmount: 500,
    Machine1CreditCount: 3,
    Machine1CreditAmount: 300,
    // ... other machine payment metrics
};