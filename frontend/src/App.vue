<template>
  <NavigationBar/>
  <router-view></router-view>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
import NavigationBar from "./components/NavigationBar.vue";

useRouter()
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
