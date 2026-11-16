<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';

</script>

<template>
  <div>
    <RouterLink to="/">Go home</RouterLink>
    <RouterLink to="/register">Sign up</RouterLink>
  </div>
  <main>
    <RouterView />
  </main>
</template>

<style scoped></style>
