<template>
  <div>
    <input v-model="text" placeholder="Enter a new item" @keyup.enter="submit" />
    <BIconArrowRightCircleFill @click="submit" />
  </div>
</template>

<script setup>
import { BIconArrowRightCircleFill } from 'bootstrap-icons-vue'
import { ref, defineEmits } from 'vue'

const text = ref('')
const emit = defineEmits(['submit'])

function submit() {
  if (text.value) {
    emit("submit", text.value);
    text.value = '';
  }
}
</script>

<style scoped>
div {
  @apply flex items-center relative;
}

input {
  @apply flex-1 rounded border border-gray-400 p-2;
}

svg {
  @apply absolute right-3 text-2xl text-blue-500 cursor-pointer transition-all;
}

svg:hover {
  @apply scale-110;
}
</style>