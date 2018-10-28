import React, { Component } from "react";
import { Modal, Radio, Slider, AutoComplete, Upload, Form, Col, Icon, Input, Button, Select } from 'antd';
import universities from "./data/universities.json";
import majors from "./data/majors.json";

const API_URL = process.env.API_URL;
const STORAGE_REQUEST_URL = 'https://www.googleapis.com/upload/storage/v1/b/tamuhack-resume-data-18/o?uploadType=media&name='

class RegistrationForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      fileName: null,
      universitySearchResult: [],
      majorSearchResult: [],
      submitModalPresent: false
    };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.onFileChange = this.onFileChange.bind(this);
    this.handleSchoolSearch = this.handleSchoolSearch.bind(this);
    this.handleMajorSearch = this.handleMajorSearch.bind(this);
  }
  
  handleSchoolSearch(query) {
    let result = universities
      .map(uni => uni.name)
      .filter(name => name.toUpperCase()
      .includes(query.toUpperCase()))
      .slice(0,5)

    this.setState({universitySearchResult: result});
  }

  handleMajorSearch(query) {
    let result = majors
      .map(major => major.Major)
      .filter(major => major.toUpperCase()
      .includes(query.toUpperCase()))
      .slice(0,5)
      .map(str => (str.charAt(0).toUpperCase() + str.substr(1).toLowerCase()));

    this.setState({majorSearchResult: result});
  }

  onFileChange(info) {
    if (info.file.status == "removed") {
      this.setState({fileName: undefined});
    } else {
      this.setState({fileName: info.file.name});
    }
  }

  handleSubmit(e) {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (err) {
        return;
      }
      const request_url = `${API_URL}` + "/registration";
      const formData = new FormData();
      formData.append("first_name", values.first_name);
      formData.append("last_name", values.last_name);
      formData.append("resume_url", values.resume_url.file.name);
      fetch(request_url, {
        method: "post",
        body: formData
      });
    });
  }

  render() {
    const { getFieldDecorator, getFieldsError, getFieldError, isFieldTouched } = this.props.form;

    const formItemLayout = {
      wrapperCol: { span: 14 },
    };

    return (
      <Form layout="horizontal" onSubmit={this.handleSubmit} style={{padding: "20px", width: "500px", maxWidth: "500px"}}>
        <p style={{fontWeight: "bold", marginBottom: "5px"}}>First Name</p>
        <Form.Item>
          {getFieldDecorator('first_name', {
            rules: [{ required: true, message: 'Please input your first name!' }],
          })(
              <Input prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />} placeholder="First Name" />
          )}
        </Form.Item>
        <p style={{fontWeight: "bold", marginBottom: "5px"}}>Last Name</p>
        <Form.Item>
          {getFieldDecorator('last_name', {
            rules: [{ required: true, message: 'Please input your last name!' }],
          })(
            <Input prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />} placeholder="Last Name" />
          )}
        </Form.Item>
        <p style={{fontWeight: "bold", marginBottom: "5px"}}>Email</p>
        <Form.Item>
          {getFieldDecorator('email', {
            rules: [{ required: true, message: 'Please input your email!' }],
          })(
            <Input prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />} placeholder="Email" />
          )}
        </Form.Item>
        <p style={{fontWeight: "bold", marginBottom: "5px"}}>Shirt Size</p>
        <Form.Item>
          {getFieldDecorator('shirt_size', {
            rules: [{ required: true, message: 'Please input your shirt size!' }],
          })(
            <Radio.Group size="medium">
              <Radio.Button value="s">Small</Radio.Button>
              <Radio.Button value="m">Medium</Radio.Button>
              <Radio.Button value="l">Large</Radio.Button>
              <Radio.Button value="xl">Extra Large</Radio.Button>
            </Radio.Group>
          )}
        </Form.Item>
        <p style={{fontWeight: "bold", marginBottom: "5px"}}>Gender</p>
        <Form.Item>
          {getFieldDecorator('gender', {
              rules: [
                { required: true, message: 'Please select your gender!' },
              ],
          })(
            <Select placeholder="Gender">
              <Select.Option value="male">Male</Select.Option>
              <Select.Option value="female">Female</Select.Option>
              <Select.Option value="other">Other</Select.Option>
              <Select.Option value="no_answer">Prefer not to answer</Select.Option>
            </Select>
          )}
        </Form.Item>
        <p style={{fontWeight: "bold", marginBottom: "5px"}}>Race</p>
        <Form.Item>
          {getFieldDecorator('race', {
              rules: [
                { required: true, message: 'Please select your gender!' },
              ],
          })(
            <Select mode="multiple" placeholder="Race">
              <Select.Option value="american_indian_alaskan_native">American Indian or Alaskan Native</Select.Option>
              <Select.Option value="asian">Asian</Select.Option>
              <Select.Option value="black_african_america">Black or African American</Select.Option>
              <Select.Option value="hispanic_latino_white">Hispanic or Latino White</Select.Option>
              <Select.Option value="native_hawaiian_pacific_islander">Native Hawaiian or other Pacific Islander</Select.Option>
              <Select.Option value="white_caucasian">White or Caucasian</Select.Option>
              <Select.Option value="no_answer">Prefer not to answer</Select.Option>
            </Select>
          )}
        </Form.Item>
        <p style={{fontWeight: "bold", marginBottom: "5px"}}>School</p>
        <Form.Item>
          {getFieldDecorator('university', {
              rules: [
                { required: true, message: 'Please select your school!' },
              ],
          })(
            <AutoComplete
              dataSource={this.state.universitySearchResult}
              onSearch={this.handleSchoolSearch}
              placeholder="University name"
            />
          )}
        </Form.Item>
        <p style={{fontWeight: "bold", marginBottom: "5px"}}>School Year</p>
        <Form.Item>
          {getFieldDecorator('school_year', {
              rules: [
                { required: true, message: 'Please select your school year!' },
              ],
          })(
            <Select placeholder="School Year">
              <Select.Option value="freshman">Freshman</Select.Option>
              <Select.Option value="sophomore">Sophomore</Select.Option>
              <Select.Option value="junior">Junior</Select.Option>
              <Select.Option value="senior">Senior</Select.Option>
              <Select.Option value="masters">Masters</Select.Option>
              <Select.Option value="phd">PhD</Select.Option>
              <Select.Option value="high_school">High Schooler</Select.Option>
            </Select>
          )}
        </Form.Item>
        <p style={{fontWeight: "bold", marginBottom: "5px"}}>Major</p>
        <Form.Item>
          {getFieldDecorator('major', {
              rules: [
                { required: true, message: 'Please select your school!' },
              ],
          })(
            <AutoComplete
              dataSource={this.state.majorSearchResult}
              onSearch={this.handleMajorSearch}
              placeholder="Major"
            />
          )}
        </Form.Item>
        <p style={{fontWeight: "bold", marginBottom: "5px"}}>How many hackathons have you been to?</p>
        <Form.Item>
          {getFieldDecorator('number_of_hackathons')(
            <Slider max={30} />
          )}
        </Form.Item>
        <p style={{fontWeight: "bold", marginBottom: "5px"}}>Why Tamuhack?</p>
        <Form.Item>
          {getFieldDecorator('why_response', {
              rules: [
                { required: true, message: 'Please input your response!' },
              ],
          })(
            <Input.TextArea
              rows={5}
            />
          )}
        </Form.Item>
        <p style={{fontWeight: "bold", marginBottom: "5px"}}>Resume</p>
        <Form.Item
          {...formItemLayout}
        >
          {getFieldDecorator('resume_url', {
              rules: [
                { required: true, message: 'Please upload your resume!' },
              ],
          })(
            <Upload 
              accept="pdf" 
              disabled={this.state.fileName != undefined && this.state.fileName != null} 
              action={STORAGE_REQUEST_URL+this.state.fileName} 
              onChange={this.onFileChange}
            >
              <Button>
                <Icon type="upload" /> Click to Upload
              </Button>
            </Upload>
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

