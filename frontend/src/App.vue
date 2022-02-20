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

<style>
:root{
  --theme-color-rgb: 139 0 0;
  --theme-color: rgb(var(--theme-color-rgb));
}

body{
  background: #eee;
  margin-block: 0;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}

p{
  line-height: 1.5;
}
</style>
