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
  --theme-color: #8B0000FF;
}

body{
  background: #ddd;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}
</style>
