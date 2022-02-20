<template>
    <div class="home">
        <Layout></Layout>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Layout from "../components/home/Layout.vue"

const name = ref<string>("Home")
</script>

<style scoped>
.home {
    width: 100%;
    height: 100%;
}
</style>