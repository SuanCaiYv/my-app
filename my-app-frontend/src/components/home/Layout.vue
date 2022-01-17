<template>
    <div class="layout">
        <slot></slot>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Layout")
</script>

<style scoped>
.layout {
    width: 100%;
    height: 100%;
    display: grid;
    grid-template-areas:
        "header"
        "body"
        "footer";
    grid-template-rows: 50px 1fr 50px;
}
</style>