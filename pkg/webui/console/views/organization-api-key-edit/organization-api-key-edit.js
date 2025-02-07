// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import PageTitle from '@ttn-lw/components/page-title'

import { ApiKeyEditForm } from '@console/components/api-key-form'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

const OrganizationApiKeyEdit = props => {
  const {
    apiKey,
    deleteOrganizationApiKey,
    deleteOrganizationApiKeySuccess,
    keyId,
    orgId,
    pseudoRights,
    rights,
    updateOrganizationApiKey,
  } = props

  useBreadcrumbs(
    'orgs.single.api-keys.edit',
    <Breadcrumb path={`/organizations/${orgId}/api-keys/${keyId}`} content={sharedMessages.edit} />,
  )

  return (
    <Container>
      <PageTitle title={sharedMessages.keyEdit} />
      <Row>
        <Col lg={8} md={12}>
          <ApiKeyEditForm
            rights={rights}
            pseudoRights={pseudoRights}
            apiKey={apiKey}
            onEdit={updateOrganizationApiKey}
            onDelete={deleteOrganizationApiKey}
            onDeleteSuccess={deleteOrganizationApiKeySuccess}
          />
        </Col>
      </Row>
    </Container>
  )
}

OrganizationApiKeyEdit.propTypes = {
  apiKey: PropTypes.apiKey.isRequired,
  deleteOrganizationApiKey: PropTypes.func.isRequired,
  deleteOrganizationApiKeySuccess: PropTypes.func.isRequired,
  keyId: PropTypes.string.isRequired,
  orgId: PropTypes.string.isRequired,
  pseudoRights: PropTypes.rights,
  rights: PropTypes.rights.isRequired,
  updateOrganizationApiKey: PropTypes.func.isRequired,
}

OrganizationApiKeyEdit.defaultProps = {
  pseudoRights: [],
}

export default OrganizationApiKeyEdit
