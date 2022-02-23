<template>
  <NavigationBar @toggleLang="toggleLang" :useEng="useEng" @toggleTheme="toggleTheme" :currentTheme="themeSetting"/>
  <router-view></router-view>
  <FooterBar/>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
import NavigationBar from "./components/NavigationBar.vue";
import FooterBar from "./components/FooterBar.vue";
import {ref} from "vue";
import {i18n} from "./plugins/translations";

useRouter()

let langFromStorage = localStorage.useEng ? JSON.parse(localStorage.useEng) : false;
let useEng = ref(langFromStorage);

function toggleLang() {
  useEng.value = !useEng.value;
  localStorage.useEng = useEng.value;
  i18n.global.locale = useEng.value ? "en-US" : "de-DE";
}

function toggleTheme() {
  let themeSetting = localStorage.theme;
  let nextTheme: { [key: string]: string; } = {
    light: "dark",
    dark: "system",
    system: "light",
  }
  //=> light => dark => system => light ....
  localStorage.theme = nextTheme[themeSetting];
  detectColorScheme()
}

let themeSetting = ref(localStorage.theme || "system");
function detectColorScheme() {
  let theme = "light";//default to light
  themeSetting.value = localStorage.theme || "system";
  if (themeSetting.value === "dark" || themeSetting.value === "system" && window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches) {
    theme = "dark";
  }
  document.documentElement.setAttribute("data-theme", theme);
  localStorage.theme = themeSetting.value;
}

detectColorScheme();

window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', () => {
  detectColorScheme();
});
</script>

<style lang="scss">
:root {
  --theme-color-rgb: 139 0 0;
  --theme-color: rgb(var(--theme-color-rgb));

  --background-color: #eee;
  --card-color: #fff;
  --text-color: black;
}

[data-theme=dark] {
  --background-color: #222;
  --card-color: #333;
  --text-color: white;
}

body {
  --footer-height: 38px;
  position: relative;
  background: var(--background-color);
  color: var(--text-color);
  margin: 0;
  padding-bottom: calc(var(--footer-height) + 10px);
  min-height: 100vh;
  box-sizing: border-box;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}

p {
  line-height: 1.5;
}

input, textarea {
  border: none;
  border-bottom: 2px solid var(--theme-color);
  background-color: var(--card-color);
  padding: 10px;
  color: inherit;
}
textarea {
  resize: vertical;
  font-family: inherit;
}

button{
  --btn-clr: var(--theme-color);
  cursor: pointer;
  position: relative;
  background:none;
  border: 2px solid var(--btn-clr);
  color: white;
  padding: 5px 10px;
  text-transform: uppercase;
  font-weight: 800;
  transition: color 0.2s ease;
  z-index: 1;

  &.light{
    --btn-clr: white;
    color: var(--theme-color);
    &:hover{
      color: var(--btn-clr);
    }
  }
  &::after{
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    background-color: var(--btn-clr);
    height:100%;
    z-index: -1;
    transition: height 0.2s ease;
  }
  &:hover{
    color: var(--btn-clr);
    &::after{
      height: 0%;
    }
  }
  &+button{
    margin-inline-start: 4px;
  }
}
</style>
