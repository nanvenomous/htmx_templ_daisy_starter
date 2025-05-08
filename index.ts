/// <reference lib="dom" />
/// <reference lib="dom.iterable" />
import htmx from "htmx.org";

const w = window as any;
w.htmx = htmx;

w.dismissToast = function(toastID: string, buttonID: string, timeout: number) {
  const toastElement = document.getElementById(toastID);
  const dismissButton = document.getElementById(buttonID);
  if (toastElement && dismissButton) {
    dismissButton.addEventListener("click", () => {
      toastElement.remove();
    });

    setTimeout(function() {
      toastElement.remove();
    }, timeout * 1000);
  }
}

const themeLocalStorageID = "theme"
const themeDefault = "corporate"
const themeDark = "aqua"
const themeDataAttribute = "data-theme"
const themeControllersSelector = "input.theme-controller"

function updateAllThemeToggles(checked: boolean) {
  const allToggles = document.querySelectorAll(themeControllersSelector);
  allToggles.forEach(toggle => {
    (toggle as HTMLInputElement).checked = checked;
  });
}

w.applySavedTheme = function() {
  const savedTheme = localStorage.getItem(themeLocalStorageID);
  document.documentElement.setAttribute(themeDataAttribute, savedTheme || themeDefault);
  var checked: boolean
  if (savedTheme === themeDark) {
    checked = true;
  } else {
    checked = false;
  }
  document.addEventListener("DOMContentLoaded", () => {
    updateAllThemeToggles(checked)
  });
}

w.persistTheme = function(checkbox: HTMLInputElement) {
  const theme = checkbox.checked ? themeDark : themeDefault;
  localStorage.setItem(themeLocalStorageID, theme);
  document.documentElement.setAttribute(themeDataAttribute, theme);
  updateAllThemeToggles(checkbox.checked)
}

