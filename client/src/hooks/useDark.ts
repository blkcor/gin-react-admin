import { darkModeAtom, toggleDarkModeAtom } from "@/stores/systemAtom";
import { useAtom, useSetAtom } from "jotai";



export const useDark = () => {
  const [darkMode] = useAtom(darkModeAtom);
  const toggleDarkMode = useSetAtom(toggleDarkModeAtom);

  return { darkMode, toggleDarkMode };
}
