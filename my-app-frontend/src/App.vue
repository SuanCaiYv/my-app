<script setup lang="ts">
import {ref} from "vue";

const name = ref<String>("app")
</script>

<template>
    <router-view></router-view>
</template>

<style>
#app {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    /* 字体大小基准 */
    font-size: 20px;
    height: 100%;
    width: 100%;
}
</style>
