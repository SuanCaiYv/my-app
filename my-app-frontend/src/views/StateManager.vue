<template>
    <div class="state-manager">
        <Layout></Layout>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import Layout from "../components/statemanager/Layout.vue"

const name = ref<String>("StateManager")
</script>

<style scoped>
.state-manager {
}
</style>