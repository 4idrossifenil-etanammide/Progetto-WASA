<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<script>
export default {
	created() {
		if(!localStorage.getItem("start")){
			localStorage.clear()
			localStorage.setItem("start", true)
		}
	}
}
</script>

<template>
	<main><RouterView /></main>
</template>

<style>
</style>
