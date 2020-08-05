import React, { ReactNode } from 'react'
import Head from 'next/head'

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
    </Head>
    <div id="valist-content">
      {children}
    </div>
    <footer>
        <hr />
        <span>I'm here to stay (Footer)</span>
    </footer>
  </div>
)

export default Layout
