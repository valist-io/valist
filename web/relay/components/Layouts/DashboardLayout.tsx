import React, { ReactNode } from 'react';
import Head from 'next/head';
import Nav from '../Navigation/NavBar';
import Footer from '../Footer/FooterBar';

type Props = {
  children?: ReactNode,
  title?: string,
};

const Layout = ({ children, title = 'Valist' }: Props): JSX.Element => (
    <div>
      <Head>
        <title>{title}</title>
        <meta charSet="utf-8" />
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      </Head>
      <div className="min-h-screen bg-gray-100">
        <Nav/>
        <main className="-mt-24 pb-8">
          <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:max-w-7xl lg:px-8">
            <h1 className="sr-only">Dashboard</h1>
            {children}
          </div>
        </main>
        <Footer />
      </div>
    </div>
);

export default Layout;
