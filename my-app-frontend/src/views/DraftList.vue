<template>
    <div class="draft-list">
        <Layout></Layout>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Layout from "../components/draftlist/Layout.vue"

const name = ref<string>("DraftList")
</script>

<style scoped>
.draft-list {
    width: 100%;
    height: 100%;
}
</style>