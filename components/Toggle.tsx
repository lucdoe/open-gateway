'use client'
import {Switch} from '@headlessui/react'
import {useState} from 'react'

function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(' ')
}

export default function Toggle() {
  const [enabled, setEnabled] = useState(false)

  return (
    <div className='mt-4'>
      <div>
        <label htmlFor='toggle' className='text-sm font-medium text-gray-900'>
          Status: {enabled ? 'Active' : 'Disabled'}
        </label>
        <div className='mt-2 '>
          <Switch
            checked={enabled}
            onChange={setEnabled}
            className='group relative inline-flex h-5 w-10 flex-shrink-0 cursor-pointer items-center justify-center rounded-full'>
            <span className='sr-only'>Use setting</span>
            <span
              aria-hidden='true'
              className='pointer-events-none absolute h-full w-full rounded-md bg-white'
            />
            <span
              aria-hidden='true'
              className={classNames(
                enabled ? 'bg-green-400' : 'bg-red-400',
                'pointer-events-none absolute mx-auto h-4 w-9 rounded-full transition-colors duration-200 ease-in-out',
              )}
            />
            <span
              aria-hidden='true'
              className={classNames(
                enabled ? 'translate-x-5' : 'translate-x-0',
                'pointer-events-none absolute left-0 inline-block h-5 w-5 transform rounded-full border border-gray-200 bg-white shadow ring-0 transition-transform duration-200 ease-in-out',
              )}
            />
          </Switch>
        </div>
      </div>
    </div>
  )
}
