import type { Alpine } from 'alpinejs';

declare global {
  interface Window {
    Alpine: Alpine;
  }
}

interface CurGetsudoRange {
    mostRecent21: string;
    closest20th: string;
}

interface CalendarDay {
    date: Date | null;
    label: number | '';
}

document.addEventListener('alpine:init', () => {

    const Alpine = window.Alpine;
    
        Alpine.data('datepicker', () => ({
            formData: {
                date:  '' as string,
                files: null as FileList | null,
            },

            innerSelectedDate: '' as string,
            currentMonth: new Date().getMonth(),
            currentYear:  new Date().getFullYear(),

            init() {
                const today = new Date();
                const year = today.getFullYear();
                const month = (today.getMonth() + 1).toString().padStart(2, '0');
                const day = today.getDate().toString().padStart(2, '0');
                this.innerSelectedDate = `${year}-${month}-${day}`;
                this.$dispatch('update-date', { data: this.innerSelectedDate });
            },

            showModal: false,

            get currentMonthYear(): string {
            return `${this.currentYear}年${this.currentMonth + 1}月`;
        },

        get calendarDays(): CalendarDay[] {
            const days: CalendarDay[] = [];
            const firstDay    = new Date(this.currentYear, this.currentMonth, 1).getDay();
            const daysInMonth = new Date(this.currentYear, this.currentMonth + 1, 0).getDate();
            for (let i = 0; i < firstDay; i++) days.push({ date: null, label: '' });
            for (let i = 1; i <= daysInMonth; i++) {
                days.push({ date: new Date(this.currentYear, this.currentMonth, i), label: i });
            }
            return days;
        },
        
        prevNextDate(offset: number): void {
            const date      = new Date(this.innerSelectedDate);
            const next_date = new Date(date.setDate(date.getDate() + offset));
            this.innerSelectedDate = `${next_date.getFullYear()}-${String(next_date.getMonth() + 1).padStart(2, '0')}-${String(next_date.getDate()).padStart(2, '0')}`;
            this.$dispatch('update-date', { data: this.innerSelectedDate });
            //this.fetchData(this.innerSelectedDate);
        },
        
        isSelected(date: Date): boolean {
            if (!this.innerSelectedDate) return false;
            const sel = new Date(this.innerSelectedDate);
            return date.getFullYear() === sel.getFullYear() &&
                   date.getMonth()    === sel.getMonth()    &&
                   date.getDate()     === sel.getDate();
        },

        isToday(date: Date): boolean {
            const today = new Date();
            return date.getFullYear() === today.getFullYear() &&
                   date.getMonth()    === today.getMonth()    &&
                   date.getDate()     === today.getDate();
        },

        selectDate(date: Date): void {
            const year  = date.getFullYear();
            const month = (date.getMonth() + 1).toString().padStart(2, '0');
            const day   = date.getDate().toString().padStart(2, '0');
            this.innerSelectedDate = `${year}-${month}-${day}`;
            this.showModal    = false;
            // this.$nextTick(() => {
            //     this.fetchData(this.innerSelectedDate);
            //     if (this.data) this.data.Record.DateString = this.innerSelectedDate;
            //     this.$dispatch('update-date', { data: this.innerSelectedDate });
            // });
        },

        previousMonth(): void {
            if (this.currentMonth === 0) { this.currentMonth = 11; this.currentYear--; }
            else this.currentMonth--;
        },

        nextMonth(): void {
            if (this.currentMonth === 11) { this.currentMonth = 0; this.currentYear++; }
            else this.currentMonth++;
        },
        
        get curGetsudoRange(): CurGetsudoRange | '' {
            if (!this.formData?.date) return '';
            const parts = String(this.formData.date).split('-');
            if (parts.length !== 3) return '';
            const y = parseInt(parts[0], 10);
            const m = parseInt(parts[1], 10) - 1;
            const d = parseInt(parts[2], 10);
            if (Number.isNaN(y) || Number.isNaN(m) || Number.isNaN(d)) return '';

            const today = new Date(y, m, d);
            const year  = today.getFullYear();
            const month = today.getMonth();
            const day   = today.getDate();

            let target: Date;
            if (day >= 21) {
                target = new Date(year, month, 21);
            } else {
                let prevMonth = month - 1;
                let yy = year;
                if (prevMonth < 0) { prevMonth = 11; yy = year - 1; }
                target = new Date(yy, prevMonth, 21);
            }
            const yyyy = target.getFullYear();
            const mm   = String(target.getMonth() + 1).padStart(2, '0');
            const dd   = String(target.getDate()).padStart(2, '0');
            const target20Date = new Date(yyyy, target.getMonth() + 1, 20);
            return {
                mostRecent21: `${yyyy}-${mm}-${dd}`,
                closest20th:  `${target20Date.getFullYear()}-${String(target20Date.getMonth() + 1).padStart(2, '0')}-${String(target20Date.getDate()).padStart(2, '0')}`,
            };
        },

        }));
    
});