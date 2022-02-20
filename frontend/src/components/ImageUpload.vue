<template>
  <input
      v-if="previewImage === ''"
      name="image"
      type="file"
      accept="image/*"
      @change="onImageUpload"
      required>
  <div v-else>
    <button @click="onImageClear">Clear</button>
    <img :src="previewImage">
  </div>

</template>

<script setup lang="ts">
import {client} from "../api/client";
import {ref} from "vue";

defineProps(['modelValue'])
const emit = defineEmits(['update:modelValue'])

const previewImage = ref('')

async function onImageUpload(event: any) {
  const resp = await client.uploadImage(event.target.files[0], true)

  previewImage.value = resp.data.thumbnailURL
  emit('update:modelValue', resp.data.id)
}

async function onImageClear(event: Event) {
  previewImage.value = ''
  emit('update:modelValue', '')
}
</script>

<style scoped lang="scss">

</style>