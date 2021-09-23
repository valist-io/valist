export const Nav = () => (
    <header className="pb-24 bg-gradient-to-r from-light-blue-800 to-violet-600">
      <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:max-w-7xl lg:px-8">
        <div className="relative flex flex-wrap items-center justify-center lg:justify-between">

          <div className="w-full py-5">
            <div className="lg:grid lg:grid-cols-3 lg:gap-8 lg:items-center">
              <div className="hidden lg:block lg:col-span-2">
                <nav className="flex space-x-4">
                  <a href="https://docs.valist.io/" className="text-violet-100 text-sm
                  font-medium rounded-md bg-white bg-opacity-0 px-3 py-2 hover:bg-opacity-10">
                    Documentation
                  </a>

                  <a href="https://valist.io/discord" className="text-violet-100 text-sm
                  font-medium rounded-md bg-white bg-opacity-0 px-3 py-2 hover:bg-opacity-10">
                    Discord
                  </a>

                  <a href="mailto:support@valist.io" className="text-violet-100 text-sm
                  font-medium rounded-md bg-white bg-opacity-0 px-3 py-2 hover:bg-opacity-10">
                    Support
                  </a>
                </nav>
              </div>
              </div>
            </div>
        </div>
      </div>
    </header>
);

export default Nav;
