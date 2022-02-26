import {Configuration, DefaultApiFactory} from "./generated";

const cfg = new Configuration({
  basePath: '/api'
});

export const client = DefaultApiFactory(cfg)
