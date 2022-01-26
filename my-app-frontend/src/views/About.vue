<template>
    <div>
        <Layout></Layout>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Layout from "../components/about/Layout.vue"

const name = ref<String>("About")
</script>

<style scoped>

</style>