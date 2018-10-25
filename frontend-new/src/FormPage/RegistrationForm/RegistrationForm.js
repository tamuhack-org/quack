import React, { Component } from "react";
import { Form, Col, Icon, Input, Button } from 'antd';

const API_URL = process.env.API_URL;

class RegistrationForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      response: "not loaded"
    };
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  hasErrors(fieldsError) {
    return Object.keys(fieldsError).some(field => fieldsError[field]);
  }

  handleSubmit(e) {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (err) {
        return;
      }
      const request_url = `${API_URL}` + "/registration";
      console.log('Received values of form: ', values);
      const formData = new FormData();
      formData.append("first_name", values.first_name);
      formData.append("last_name", values.last_name);
      // Set a loading state before doing anything else.
      this.setState({ loadingState: "loading" }, () => {
        // Send a post to our backend.
        fetch(request_url, {
          method: "post",
          body: formData
          // Regardless of the speed of the request, we wait atleast 1.5 seconds.
        }).then(() => {
          setTimeout(() => {
            this.setState({ loadingState: "done" });
          }, 1500);
        });
      });
    });
  }

  render() {
    const { getFieldDecorator, getFieldsError, getFieldError, isFieldTouched } = this.props.form;

    const firstNameError = getFieldError('first_name');
    const lastNameError = getFieldError('last_name');

    const formItemLayout = {
      wrapperCol: { span: 14 },
    };

    return (
      <Form layout="horizontal" onSubmit={this.handleSubmit} style={{width: "300px"}}>
        <Form.Item
          validateStatus={firstNameError ? 'error' : ''}
          help={firstNameError || ''}
       >
          {getFieldDecorator('first_name', {
            rules: [{ required: true, message: 'Please input your first name!' }],
          })(
              <Input prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />} placeholder="First Name" />
          )}
        </Form.Item>
        <Form.Item
          validateStatus={lastNameError ? 'error' : ''}
          help={lastNameError || ''}
        >
          {getFieldDecorator('last_name', {
            rules: [{ required: true, message: 'Please input your last name!' }],
          })(
            <Input prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />} placeholder="Last Name" />
          )}
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">Submit</Button>
        </Form.Item>
      </Form>
    );
  }
}

export default Form.create()(RegistrationForm);

