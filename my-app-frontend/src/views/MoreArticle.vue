<template>
    <div class="more-article">
        <Layout></Layout>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Layout from "../components/morearticle/Layout.vue"

const name = ref<String>("MoreArticle")
</script>

<style scoped>
.more-article {
    width: 100%;
    height: 100%;
}
</style>