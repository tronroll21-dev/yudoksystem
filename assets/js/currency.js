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

    // When the element gains focus, remove commas for editing
    const handleFocus = (e) => {
      e.target.value = getNumericValue(e.target.value);
    };

    // When the element loses focus, update the Alpine model and re-format
    const handleBlur = (e) => {
      const numericValue = getNumericValue(e.target.value);
      const valueToSet = numericValue === '' ? 0 : parseInt(numericValue, 10);
      setValue(valueToSet);

      const formattedValue = numericValue === '' ? '' : formatWithCommas(numericValue);
      if (e.target.value !== formattedValue) {
        e.target.value = formattedValue;
      }
    };

    // Watch for programmatic changes to the model and update the display
    effect(() => {
      getValue((value) => {
        // Do not format if the element is focused, as the user is editing.
        if (document.activeElement === el) {
          return;
        }

        const numValue = value && !isNaN(value) ? parseInt(value, 10) : 0;
        const formattedValue = numValue === 0 ? '' : formatWithCommas(numValue);
        
        if (el.value !== formattedValue) {
          el.value = formattedValue;
        }
      });
    });

    // Attach event listeners
    el.addEventListener('focus', handleFocus);
    el.addEventListener('blur', handleBlur);

    // Cleanup
    cleanup(() => {
      el.removeEventListener('focus', handleFocus);
      el.removeEventListener('blur', handleBlur);
    });
  });
});