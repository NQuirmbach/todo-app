<template>
  <ul v-if="props.items.length > 0">
    <li v-for="item in props.items" v-bind:key="item.id">
      <input type="checkbox" :value="item.isDone" @input="e => item.isDone = e.target.checked" />
      <input type="text" :value="item.text" @input="e => item.text = e.target.value" />
      <BIconTrash @click="() => emit('remove', item)" />
    </li>
  </ul>

  <div v-if="props.items.length === 0" class="italic text-center">No items...</div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue'
import { BIconTrash } from 'bootstrap-icons-vue'

const props = defineProps({
  items: Array
})
const emit = defineEmits('remove')
</script>

<style scoped lang="scss">
li {
  @apply flex rounded border border-gray-400 py-2 px-4 transition-shadow gap-2 items-center mb-4;
}

li:hover {
  @apply shadow;
}

input[type="text"] {
  @apply flex-1 outline-none;
}

svg {
  @apply cursor-pointer transition-colors;

  &:hover {
    @apply text-red-400;
  }
}
</style>