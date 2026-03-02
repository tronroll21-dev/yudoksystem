document.addEventListener('alpine:init', () => {
Alpine.data('uploader', () => ({
        formData: {
            date: '',
            files: null,
        },       
        updateDate(newDate) {
            console.log('Updating date to:', newDate);
            this.formData.date = newDate;
        },
        get curCalendarMonth() {
            if (!this.formData || !this.formData.date) return '';
            const parts = String(this.formData.date).split('-');
            if (parts.length !== 3) return '';
            const y = parseInt(parts[0], 10);
            const m = parseInt(parts[1], 10) - 1;
            const d = parseInt(parts[2], 10);
            if (Number.isNaN(y) || Number.isNaN(m) || Number.isNaN(d)) return '';   
            const today = new Date(y, m, d);
            const year = today.getFullYear();
            const month = today.getMonth();
            const firstDay = new Date(year, month, 1);
            const lastDay = new Date(year, month + 1, 0);
            const yyyy = firstDay.getFullYear();
            const mm = String(firstDay.getMonth() + 1).padStart(2, '0');
            const dd1 = String(firstDay.getDate()).padStart(2, '0');
            const dd2 = String(lastDay.getDate()).padStart(2, '0');
            return {startDate:`${yyyy}-${mm}-${dd1}`, endDate:`${yyyy}-${mm}-${dd2}`};
        },
        get curGetsudoRange() {
            if (!this.formData || !this.formData.date) return '';
            const parts = String(this.formData.date).split('-');
            if (parts.length !== 3) return '';
            const y = parseInt(parts[0], 10);
            const m = parseInt(parts[1], 10) - 1; // 0-based month
            const d = parseInt(parts[2], 10);
            if (Number.isNaN(y) || Number.isNaN(m) || Number.isNaN(d)) return '';

            const today = new Date(y, m, d);
            const year = today.getFullYear();
            const month = today.getMonth(); // 0-based
            const day = today.getDate();
            let target;
            if (day >= 21) {
                target = new Date(year, month, 21);
            } else {
                let prevMonth = month - 1;
                let yy = year;
                if (prevMonth < 0) { prevMonth = 11; yy = year - 1; }
                target = new Date(yy, prevMonth, 21);
            }
            const yyyy = target.getFullYear();
            const mm = String(target.getMonth() + 1).padStart(2, '0');
            const dd = String(target.getDate()).padStart(2, '0');
            const target20Date = new Date(yyyy, target.getMonth() + 1, 20);
            return {mostRecent21:`${yyyy}-${mm}-${dd}`, closest20th: target20Date.getFullYear() + '-' + String(target20Date.getMonth() + 1).padStart(2, '0') + '-' + String(target20Date.getDate()).padStart(2, '0')    };
        }
     

    }))
});