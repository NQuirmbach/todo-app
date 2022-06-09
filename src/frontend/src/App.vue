<template>
  <h1 class="text-xl mb-2">My ToDos</h1>
  <NewTodo @submit="submitNewItem" />

  <hr />
  <TodoList v-if="!loading" :items="items" @remove="removeItem" />
</template>

<script setup>
import axios from 'axios';
import NewTodo from './components/NewTodo.vue'
import TodoList from './components/TodoList.vue'

import { onMounted, ref } from 'vue'

const items = ref([]);
const loading = ref(true);

async function submitNewItem(text) {
  const res = await axios.post('/api/todos', {
    text,
    isDone: false,
  });
  items.value.push(res.data)
}
async function removeItem(item) {
  await axios.delete(`/api/todos/${item.id}`);
  items.value.splice(items.value.indexOf(item), 1);
}

onMounted(async () => {
  try {
    const res = await axios.get('/api/todos')
    items.value = res.data;
  } catch (err) {
    console.error(err);
  }
  loading.value = false;
})
</script>

<style>
@tailwind base;
@tailwind components;
@tailwind utilities;

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;

  @apply w-3/6 mx-auto mt-14;
}

h1 {
  @apply text-center;
}

hr {
  @apply my-4;
}
</style>
