import { HellosClient } from "../../js/hello_pb_service.js";
import { HelloRequest, HelloStreamRequest } from "../../js/hello_pb.js";

export default {
  install(Vue) {
    Vue.prototype.client = new HellosClient(process.env.VUE_APP_BASE_URL, {
      withCredentials: true,
    });

    Vue.prototype.sayHello = function (t) {
      return new Promise((resolve, reject) => {
        try {
          const req = new HelloRequest();
          req.setRequest(t);
          Vue.prototype.client.sayHello(req, {}, (err, resp) => {
            if (err != null) {
              reject(new Error(JSON.stringify(err)));
              return false;
            }
            if (resp == null) {
              reject(new Error("no resp"));
              return false;
            }
            resolve(resp.getResponse());
          });
        } catch (err) {
          reject(err);
        }
      });
    };
    Vue.prototype.helloStream = function (auth) {
      const req = new HelloStreamRequest();
      req.setAuth(auth);
      return Vue.prototype.client.helloStream(req, {});
    };
  },
};
