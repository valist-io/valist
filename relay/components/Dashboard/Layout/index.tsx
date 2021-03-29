import React, { ReactNode } from 'react';
import Head from 'next/head';
import Nav from '../Nav';
import Footer from '../Footer';

type Props = {
  children?: ReactNode,
  title?: string,
}

const Layout = ({ children, title = 'Valist' }: Props) => {
  return (
    <div>
      <Head>
        <title>{title}</title>
        <meta charSet="utf-8" />
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      </Head>
      <div className="min-h-screen bg-gray-100">
        <Nav/>
        {children}
        <Footer />
      </div>
    </div>
  )
}

export default Layout;
