<template>
  <ol id="commentList">
    <li class="comment" v-for="comment in comments">
      <div class="meta">
        <span class="comment--author">{{ comment.name }} </span>
        <span class="comment--date"> {{ comment.date }}</span>
        <button v-if="useAuth.loggedIn">Delete</button>
      </div>
      <p class="comment--text">{{ comment.text }}</p>
    </li>
  </ol>
</template>
<script lang="ts" setup>
import {useAuth} from "../stores/auth";

defineProps<{
  comments: []
}>()
</script>
<style scoped lang="scss">

#commentList {
  text-align: start;
  --padding: 20px;
  list-style: none;
  padding: 0;
}

.comment {
  overflow: auto;
  margin-inline: calc(-1 * var(--padding));
  padding-inline: calc(var(--padding));
  padding-top: var(--padding);

  &:nth-child(even) {
    background-color: rgba(var(--theme-color-rgb) / 0.2);
  }

  &--author, &--date {
    opacity: 0.7;
  }

  &--author {
    margin-right: auto;
  }

  .meta{
    display: flex;
    flex-direction: row;
    align-items: center;
  }
}
</style>
