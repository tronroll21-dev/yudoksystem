// Alpine.js Currency Directive
document.addEventListener('alpine:init', () => {
  Alpine.directive('currency', (el, { expression }, { evaluateLater, effect, evaluate, cleanup }) => {
    // Ensure the element is an input
    if (el.tagName !== 'INPUT') {
      console.warn('x-currency directive should only be used on input elements');
      return;
    }

    const getValue = evaluateLater(expression);
    const setValue = (value) => {
      evaluate(`${expression} = ${value}`);
    };

    // Format number with commas
    const formatWithCommas = (num) => {
      return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',');
    };

    // Remove all non-digit characters
    const getNumericValue = (str) => {
      return str.replace(/\D/g, '');
    };

    // Handle input event
    const handleInput = (e) => {
      const cursorPosition = e.target.selectionStart;
      const oldValue = e.target.value;
      const oldLength = oldValue.length;

      // Get only numeric characters
      const numericValue = getNumericValue(oldValue);

      // Update the Alpine data with numeric value
      setValue(numericValue === '' ? 0 : parseInt(numericValue, 10));

      // Format with commas for display
      const formattedValue = numericValue === '' ? '' : formatWithCommas(numericValue);
      e.target.value = formattedValue;

      // Adjust cursor position after formatting
      const newLength = formattedValue.length;
      const diff = newLength - oldLength;
      const newPosition = cursorPosition + diff;
      
      // Set cursor position
      requestAnimationFrame(() => {
        e.target.setSelectionRange(newPosition, newPosition);
      });
    };

    // Watch for programmatic changes and update display
    effect(() => {
      getValue((value) => {
        const numValue = value && !isNaN(value) ? parseInt(value, 10) : 0;
        const formattedValue = numValue === 0 ? '' : formatWithCommas(numValue);
        
        // Only update if the formatted value differs from current display
        // This prevents cursor jumping during manual input
        if (el.value !== formattedValue) {
          el.value = formattedValue;
        }
      });
    });

    // Attach input event listener
    el.addEventListener('input', handleInput);

    // Cleanup
    cleanup(() => {
      el.removeEventListener('input', handleInput);
    });
  });
});