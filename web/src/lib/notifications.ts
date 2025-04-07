export function showNotification(type: 'success' | 'error', message: string) {
  const notification = document.createElement('div');
  notification.innerHTML = `
    <div class="fixed bottom-4 right-4 z-50">
      <div class="p-4 rounded-lg shadow-lg ${
        type === 'success' ? 'bg-green-500' : 'bg-red-500'
      } text-white">
        ${message}
      </div>
    </div>
  `;
  document.body.appendChild(notification);

  // Auto-hide after 2 seconds
  setTimeout(() => {
    notification.remove();
  }, 2000);
}

export function getErrorMessage(status: number): string {
  switch (status) {
    case 400:
      return 'Invalid request. Please check your input and try again.';
    case 401:
      return 'Unauthorized. Please check your credentials.';
    case 403:
      return 'Forbidden. You do not have permission to perform this action.';
    case 404:
      return 'Resource not found. Please check the URL and try again.';
    case 409:
      return 'Conflict. This feed might already exist.';
    case 500:
      return 'Server error. Please try again later.';
    default:
      return 'An unexpected error occurred. Please try again.';
  }
} 