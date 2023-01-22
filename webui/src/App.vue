<script setup>
import { RouterLink, RouterView } from 'vue-router'
import HomeView from './views/HomeView.vue';
</script>

<script>
localStorage.clear();
export default {
	components: {
		HomeView
	}
}
</script>

<template>
	<main><RouterView /></main>
</template>

<style>
</style>
