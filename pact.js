import { pactWith } from 'jest-pact';
import { like, eachLike, string } from "@pact-foundation/pact/src/dsl/matchers";
import { API } from "@/api";

pactWith({
    consumer: 'FrontEnd',
    provider: 'Backend',

}, provider => {
    describe('animal diseases', () => {
        let api
        beforeEach(() => {
            console.log(provider.mockService.baseUrl)
            api = new API(provider.mockService.baseUrl)
        })

        it('get all animals correctly', async function () {
            await provider.addInteraction({
                state: 'when user enters the animals page',
                uponReceiving: 'a request for animal list',
                withRequest: {
                    method: 'GET',
                    path: '/animals'
                },
                willRespondWith: {
                    status: 200,
                    headers: {
                        'Content-Type': 'application/json; charset=UTF-8',
                    },
                    body: {
                        data: eachLike({
                            id: like(1),
                            name: string('Dog'),
                            type: string(''),
                            order: like(1),
                            image: string('')
                        })
                    }
                }
            })
            const response = await api.getAnimals()
            expect(response.data[0].id).toEqual(1)
            expect(response.data[0].name).toEqual('Dog')

        })

        it('get all symptoms based on animal correctly', async function () {
            await provider.addInteraction({
                state: 'when user selects an animal',
                uponReceiving: 'a request for all symptoms',
                withRequest: {
                    method: 'GET',
                    path: '/animal/:id/symptoms',
                    body: {
                    }
                },
                willRespondWith: {
                    status: 200,
                    headers: {
                        'Content-Type': 'application/json; charset=UTF-8',
                    },
                    body: {
                        data: {
                            id: like(1),
                            animal_id: like(1),
                            desc: string(''),
                            created_at: iso8601Date('2016-01-01'),
                            updated_at: iso8601Date('2017-01-01'),
                            initial_action_id: integer(132)
                        }
                    }
                }
            })

            const response = await api.getSymptom()
            expect(response.data.id).toEqual(1)
            expect(response.data.animal_id).toEqual(1)
            expect(response.data.initial_action_id).toEqual(132)

        })

        it('get actionable', async function () {
            await provider.addInteraction({
                state: 'when user selects a symptom',
                uponReceiving: 'a request for get actionables',
                withRequest: {
                    method: 'GET',
                    path: '/sypmtom/:id'
                },
                willRespondWith: {
                    status: 200,
                    headers: {
                        'Content-Type': 'application/json; charset=UTF-8',
                    },
                    body: {
                        data: eachLike({
                            id: like(1),
                            created_at: iso8601Date('2016-01-01'),
                            updated_at: iso8601Date('2017-01-01'),
                            type: string(''),
                            question: eachLike({
                                actionable_id: integer,
                                text: string(''),
                                options: eachLike({
                                    response_id: integer,
                                    text: string(''),
                                    next_action_id: integer
                                })
                            })
                        })
                    }
                }
            })
            const response = await api.getActionable()
            expect(response.data[0].id).toEqual(1)
            // todo: more test cases

        })
        it('get result', async function () {
            await provider.addInteraction({
                state: 'when user selects an option & result will be shown ',
                uponReceiving: 'a request for get final result',
                withRequest: {
                    method: 'GET',
                    path: '/result/:id'
                },
                willRespondWith: {
                    status: 200,
                    headers: {
                        'Content-Type': 'application/json; charset=UTF-8',
                    },
                    body: {
                        data: {
                            response_id: like(1),
                            additional_advice: string(''),
                            first_aid_text: string(''),
                            problem_text: string(''),
                            travel_advice_text: string(''),
                            iframe_first_aid_text: string(''),
                            iframe_problem_text: string(''),
                            risk_category_id: like(1)
                        }
                    }
                }
            })
            const response = await api.getActionable()
            expect(response.data[0].id).toEqual(1)
            // todo: more test cases

        })
    })
})
