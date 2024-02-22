'use client'
import TableWithHeaderAndAddEntry from '@/components/TableWithHeaderAndAddEntry'
import {apiRoutes} from '@/config/config'
import {useEffect, useState} from 'react'

export const endpointsFields = [
  {fieldtype: 'toggle', label: 'Enabled', value: 'true'},
  {fieldtype: 'text', label: 'Name', placeholder: 'Test Name'},

  {fieldtype: 'text', label: 'Path', placeholder: '/test'},
  {
    fieldtype: 'checkbox',
    label: 'Methods',
    name: 'Methods',
    fields: ['GET', 'POST', 'PUT', 'DELETE', 'PATCH', 'OPTIONS'],
  },
  {
    fieldtype: 'text',
    label: 'Tags',
    placeholder: 'test, example, tag',
  },
]

export default function Endpoints() {
  const [endpoints, setEndpoints] = useState(null)

  useEffect(() => {
    const fetchData = async () => {
      const res = await fetch(apiRoutes.endpoints.all)
      const data = await res.json()
      setEndpoints(data)
    }

    fetchData()
  }, [])

  return (
    <div>
      <TableWithHeaderAndAddEntry
        name='Endpoints'
        description='List of all endpoints, currently added to your cluster.'
        tableRows={endpoints ?? []}
      />
    </div>
  )
}
