<template>
    <div class="sign">
        <Layout></Layout>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Layout from "../components/sign/Layout.vue";

const name = ref<String>("Sign")
</script>

<style scoped>
.sign {
    width: 100%;
    height: 100%;
}
</style>