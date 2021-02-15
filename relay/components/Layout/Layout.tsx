import React, { ReactNode } from 'react'
import Head from 'next/head'
import Nav from '../Nav/Nav';

type Props = {
  children?: ReactNode,
  title?: string,
}

const Layout = ({ children, title = 'Valist' }: Props) => (

  <div>
    <Head>
      <title>{title}</title>
      <meta charSet="utf-8" />
      <meta name="viewport" content="initial-scale=1.0, width=device-width" />
    </Head>
    <Nav />
    <div id="valist-content-fixed">
      {children}
    </div>
  </div>
)

export default Layout
