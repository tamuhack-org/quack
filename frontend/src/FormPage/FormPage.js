import React, { Component } from "react";
import RegistrationForm from "./RegistrationForm/RegistrationForm.js";
import { FormWrapper } from "./FormPageStyles.js";
import { Flex, Box } from "grid-styled";

class FormPage extends Component {
  render() {
    return (
      <Flex alignItems="center" justifyContent="center"
          css={{
            width: "100vw",
            minHeight: "100vh",
          }}
      >
        <RegistrationForm/>
      </Flex>
    );
  }
}

export default FormPage;
