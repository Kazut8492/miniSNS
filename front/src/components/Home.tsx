import React from "react";
import {useEffect, useState} from "react";
import {Breadcrumb, Layout, Menu, theme} from "antd";

const {Header, Content, Sider} = Layout;

export const Home: React.FC<{  }> = props => {
    const [data, setData] =  useState<any[]>([])

    useEffect(() => {
        fetch("http://localhost:8080/home")
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setData(data);
            }
        );
    }, []);

    return (
        <Layout>
            <Header>
                <div className="logo" />
                <Menu />
            </Header>
            <Layout>
                <Sider>
                    <Menu />
                </Sider>
                <Layout>
                    <Breadcrumb />
                    <Content>Content</Content>
                </Layout>
            </Layout>
        </Layout>
    );
};