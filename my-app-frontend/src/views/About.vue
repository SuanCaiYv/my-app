<template>
    <div class="about">
        <Layout></Layout>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Layout from "../components/about/Layout.vue"

const name = ref<String>("About")
</script>

<style scoped>
.about {
    width: 100%;
    height: 100%;
}
</style>