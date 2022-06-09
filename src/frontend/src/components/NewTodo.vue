<template>
  <input v-model="value" placeholder="Enter a new item" />

</template>

<script setup>
import { ref } from 'vue'

const value = ref('')
</script>

<style scoped>
input {
  @apply rounded border border-gray-400 p-2 w-full;
}
</style>