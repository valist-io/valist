import React, { useContext, useEffect, useState } from 'react';
import Layout from '../components/Layout/Layout';

export const SettingsPage = () => {

    return (
        <Layout title="Valist | Create Organization">
            <div>
                <h1 style={{fontSize: 45, textAlign: "center"}}>Local Device Settings</h1>
            </div>
        </Layout>
    );
}

export default SettingsPage;
