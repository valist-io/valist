import React, { ReactNode } from 'react'
import Head from 'next/head'
import Nav from '../Nav/NavTop';

type Props = {
  children?: ReactNode
  title?: string
}

const Layout = ({ children, title = 'Valist' }: Props) => (

  <div>
    <Head>
      <title>{title}</title>
      <meta charSet="utf-8" />
      <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      <link href="https://fonts.googleapis.com/css2?family=Noto+Sans:wght@400;700&display=swap" rel="stylesheet" />
    </Head>
    <Nav />
    <div id="valist-content-fixed">
      {children}
    </div>
  </div>
)

export default Layout
