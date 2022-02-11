<template>
    <div class="editor">
        <Layout></Layout>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Layout from "../components/editor/Layout.vue"

const name = ref<String>("Editor")
</script>

<style scoped>
.editor {
    width: 100%;
    height: 100%;
}
</style>