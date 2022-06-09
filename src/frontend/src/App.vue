<template>
  <h1 class="text-xl mb-2">ToDo App</h1>
  <NewTodo />
  <HelloWorld msg="Welcome to Your Vue.js App" />
</template>

<script>
import NewTodo from './components/NewTodo.vue'

export default {
  name: 'App',
  components: {
    NewTodo,
  }
}
</script>

<style lang="scss">
@tailwind base;
@tailwind components;
@tailwind utilities;

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;

  @apply w-3/6 mx-auto;
}
</style>
