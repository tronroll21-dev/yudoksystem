import type { Alpine } from 'alpinejs';

declare global {
  interface Window {
    Alpine: Alpine;
    Chart: any;
    Temporal: any;
  }
}

interface AnalysisData {
    date: string;
    hotbath_revenue: number;
    restaurant_revenue: number;
    relaxation_revenue: number;
    power_expense: number;
    gas_expense: number;
    visitor_count: number;
}

document.addEventListener('alpine:init', () => {
    window.Alpine.data('analysisPage', () => {
        // Moving the chart instance to a non-reactive closure variable.
        // This prevents Alpine.js from attempting to make the complex Chart object reactive,
        // which causes the "Maximum call stack size exceeded" error.
        let chartInstance: any = null;

        return {
            data: [] as AnalysisData[],
            isLoading: false,
            startDate: '',
            endDate: '',

            async init() {
                const range = this.calculateInitialDates();
                this.startDate = range.start;
                this.endDate = range.end;

                await this.fetchData();
            },

            isValidDate(dateStr: string): boolean {
                if (!dateStr) return false;
                if (!/^\d{4}-\d{2}-\d{2}$/.test(dateStr)) return false;
                
                try {
                    const parts = dateStr.split('-');
                    const year = parseInt(parts[0], 10);
                    if (year < 2000 || year > 2100) return false;

                    window.Temporal.PlainDate.from(dateStr);
                    return true;
                } catch (e) {
                    return false;
                }
            },

            calculateInitialDates() {
                const Temporal = window.Temporal;
                const today = Temporal.Now.plainDateISO();
                let end: any;
                
                if (today.day < 20) {
                    end = today.subtract({ months: 1 }).with({ day: 20 });
                } else {
                    end = today.with({ day: 20 });
                }
                
                const start = end.subtract({ months: 1 }).with({ day: 21 });
                
                return {
                    start: start.toString(),
                    end: end.toString()
                };
            },

            async fetchData() {
                if (!this.isValidDate(this.startDate) || !this.isValidDate(this.endDate)) {
                    return;
                }

                this.isLoading = true;
                try {
                    const res = await fetch(`/api/analysis-data?start_date=${this.startDate}&end_date=${this.endDate}`);
                    if (!res.ok) throw new Error('Failed to fetch analysis data');
                    this.data = await res.json();
                    
                    // Re-initialize to avoid Chart.js internal scale errors (like the "fullSize" bug)
                    this.initChart();
                } catch (err) {
                    console.error(err);
                    alert('データの取得に失敗しました。');
                } finally {
                    this.isLoading = false;
                }
            },

            initChart() {
                const ctx = (document.getElementById('analysisChart') as HTMLCanvasElement)?.getContext('2d');
                if (!ctx) return;

                // Create a clean copy of the data, stripping any Alpine proxies to prevent Chart.js recursion bugs.
                const rawData = JSON.parse(JSON.stringify(this.data));
                const labels = rawData.map((d: any) => d.date.split('-').slice(1).join('/'));
                
                const datasets = [
                    {
                        label: '入浴売上',
                        data: rawData.map((d: any) => d.hotbath_revenue),
                        borderColor: 'rgb(54, 162, 235)',
                        backgroundColor: 'rgba(54, 162, 235, 0.5)',
                        yAxisID: 'y',
                        tension: 0.1
                    },
                    {
                        label: '飲食売上',
                        data: rawData.map((d: any) => d.restaurant_revenue),
                        borderColor: 'rgb(255, 99, 132)',
                        backgroundColor: 'rgba(255, 99, 132, 0.5)',
                        yAxisID: 'y',
                        tension: 0.1
                    },
                    {
                        label: 'リラク売上',
                        data: rawData.map((d: any) => d.relaxation_revenue),
                        borderColor: 'rgb(75, 192, 192)',
                        backgroundColor: 'rgba(75, 192, 192, 0.5)',
                        yAxisID: 'y',
                        tension: 0.1
                    },
                    {
                        label: '電気代',
                        data: rawData.map((d: any) => d.power_expense),
                        borderColor: 'rgb(255, 205, 86)',
                        backgroundColor: 'rgba(255, 205, 86, 0.5)',
                        yAxisID: 'y',
                        tension: 0.1
                    },
                    {
                        label: 'ガス代',
                        data: rawData.map((d: any) => d.gas_expense),
                        borderColor: 'rgb(255, 159, 64)',
                        backgroundColor: 'rgba(255, 159, 64, 0.5)',
                        yAxisID: 'y',
                        tension: 0.1
                    },
                    {
                        label: '来客数',
                        data: rawData.map((d: any) => d.visitor_count),
                        borderColor: 'rgb(153, 102, 255)',
                        backgroundColor: 'rgba(153, 102, 255, 0.5)',
                        yAxisID: 'y1',
                        type: 'bar' as const,
                    }
                ];

                if (chartInstance) {
                    chartInstance.destroy();
                }

                chartInstance = new window.Chart(ctx, {
                    type: 'line',
                    data: {
                        labels: labels,
                        datasets: datasets
                    },
                    options: {
                        responsive: true,
                        animation: {
                            duration: 0 // Smooth the transition during re-creation
                        },
                        interaction: {
                            mode: 'index',
                            intersect: false,
                        },
                        scales: {
                            y: {
                                type: 'linear',
                                display: true,
                                position: 'left',
                                title: {
                                    display: true,
                                    text: '金額 (円)'
                                }
                            },
                            y1: {
                                type: 'linear',
                                display: true,
                                position: 'right',
                                grid: {
                                    drawOnChartArea: false,
                                },
                                title: {
                                    display: true,
                                    text: '来客数'
                                }
                            },
                        },
                        plugins: {
                            zoom: {
                                pan: {
                                    enabled: true,
                                    mode: 'x',
                                },
                                zoom: {
                                    drag: {
                                        enabled: true,
                                    },
                                    mode: 'x',
                                }
                            }
                        }
                    }
                });

                // Add double-click handler to navigate to report
                ctx.canvas.ondblclick = (evt) => {
                    const points = chartInstance.getElementsAtEventForMode(evt, 'nearest', { intersect: true }, true);
                    if (points.length) {
                        const index = points[0].index;
                        const date = this.data[index].date;
                        window.open(`/api/sales-data-report-svg?date=${date}`, '_blank');
                    }
                };
            },

            resetZoom() {
                if (chartInstance) {
                    chartInstance.resetZoom();
                }
            }
        };
    });
});
