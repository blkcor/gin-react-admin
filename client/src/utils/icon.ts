import { loadIcons } from '@iconify/react';

export function loadIconAsync(icons: string[]) {
  return new Promise((fulfill, reject) => {
    loadIcons(icons, (loaded, missing, pending) => {
      if (pending.length) {
        return
      }
      if (missing.length) {
        reject({
          loaded,
          missing,
        });
      } else {
        fulfill({
          loaded,
        });
      }
    })
  })
}
