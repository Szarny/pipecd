import { Toasts } from "./";
import { Provider } from "react-redux";
import { createStore } from "test-utils";
import { Story } from "@storybook/react";

export default {
  title: "COMMON/Toasts",
  component: Toasts,
};

export const Overview: Story = () => (
  <Provider
    store={createStore({
      toasts: {
        entities: {
          "1": {
            id: "1",
            message: "toast message",
            severity: "error",
          },
        },
        ids: ["1"],
      },
    })}
  >
    <Toasts />
  </Provider>
);

export const Url: Story = () => (
  <Provider
    store={createStore({
      toasts: {
        entities: {
          "1": {
            id: "1",
            message: "toast message",
            severity: "success",
            to: "/",
          },
        },
        ids: ["1"],
      },
    })}
  >
    <Toasts />
  </Provider>
);
