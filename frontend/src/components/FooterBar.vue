<template>
  <FindUs/>
  <footer>
    <a v-if="!auth.loggedIn" href="/auth/login">{{ $t('navigation.login') }}</a>
    <a v-else href="/auth/logout">{{ $t('navigation.logout') }}</a>
    <div v-if="auth.loggedIn" id="activeUser">Logged in as: <span>{{auth.user}}</span></div>
    <router-link to="/legal">{{ $t('navigation.imprint') }}</router-link>
  </footer>
</template>

<script setup lang="ts">
import {useAuth} from "../stores/auth";
import FindUs from "./FindUs.vue";

const auth = useAuth()
</script>

<style scoped lang="scss">
footer{
  height: var(--footerHeight);
  box-sizing: border-box;
  position: absolute;
  bottom: 0;
  width: 100%;
  margin-top: 10px;
  padding: 10px 20px;
  background-color: #111;
  font-weight: bold;
  display: flex;
  justify-content: space-between;

  #activeUser{
    font-weight: initial;
  }
  a{
    text-decoration: none;
    color: white;

    &:hover{
      text-decoration: underline;
    }
  }
}
</style>
