<template>
  <footer>
    <a v-if="!auth.loggedIn" href="/auth/login">{{ $t('navigation.login') }}</a>
    <a v-else href="/auth/logout">{{ $t('navigation.logout') }}</a>
    <router-link to="/legal">{{ $t('navigation.imprint') }}</router-link>
  </footer>
</template>

<script setup lang="ts">
import {useAuth} from "../stores/auth";

const auth = useAuth()
</script>

<style scoped lang="scss">
footer{
  margin-top: 10px;
  padding: 10px 20px;
  background-color: #111;
  font-weight: bold;
  display: flex;
  justify-content: space-between;
  a{
    text-decoration: none;
    color: white;

    &:hover{
      text-decoration: underline;
    }
  }
}
</style>
