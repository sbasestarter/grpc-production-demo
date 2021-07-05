<template>
  <div id="app">
    <el-button @click="langSwitch">el-button</el-button>
    <p>{{ $t("test") }}</p>
    <el-input type="textarea" :rows="20" :placeholder="testText()" />
    <el-button @click="doHelloStream">helloStream</el-button>
  </div>
</template>

<script>
export default {
  name: "app",
  methods: {
    testText() {
      return this.$i18n.t("test");
    },
    langSwitch() {
      this.sayHello("test??").then((resp) => {
        console.log(resp);
      });
      if (this.$i18n.locale === "zh_CN") {
        this.$i18n.locale = "en_US";
      } else {
        this.$i18n.locale = "zh_CN";
      }
    },
    doHelloStream() {
      this.helloStream("123957")
        .on("data", (msg) => {
          console.log(msg.getMessage());
        })
        .on("end", (code, details, metadata) => {
          console.log(code, details, metadata);
        });
    },
  },
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
