/// <reference lib="dom" />
/// <reference lib="dom.iterable" />
import htmx from "htmx.org";

const w = window as any;
w.htmx = htmx;

w.dismissToast = function (toastID: string, buttonID: string, timeout: number) {
  const toastElement = document.getElementById(toastID);
  const dismissButton = document.getElementById(buttonID);
  if (toastElement && dismissButton) {
    dismissButton.addEventListener("click", () => {
      toastElement.remove();
    });

    setTimeout(function () {
      toastElement.remove();
    }, timeout * 1000);
  }
}
