<template>
  <NavigationBar/>
  <router-view></router-view>
  <FooterBar/>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
import NavigationBar from "./components/NavigationBar.vue";
import FooterBar from "./components/FooterBar.vue";

useRouter()
</script>

<style lang="scss">
:root {
  --theme-color-rgb: 139 0 0;
  --theme-color: rgb(var(--theme-color-rgb));
}

body {
  --footerHeight: 38px;
  position: relative;
  background: #eee;
  margin: 0;
  padding-bottom: calc(var(--footerHeight) + 10px);
}

@media (prefers-color-scheme: dark) {
  :root {
    --card-color: #333
  }

  body{
    background: #222;
    color: white;
  }
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
</style>
